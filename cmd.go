package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const rtagFile = ".rtag"

var rootCmd *cobra.Command
var initCmd *cobra.Command
var addCmd *cobra.Command
var pushCmd *cobra.Command
var listCmd *cobra.Command
var rmCmd *cobra.Command

var pushAll bool

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing root command: %v", err)
		os.Exit(1)
	}
}

func init() {
	initializeCommands()
}

func initializeCommands() {
	// Initialize i18n
	InitI18n()

	// Initialize root command with translated strings
	rootCmd = &cobra.Command{
		Use:   "rtag",
		Short: T().RootShort,
		Long:  T().RootLong,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	initCmd = &cobra.Command{
		Use:   "init",
		Short: T().InitShort,
		Long:  T().InitLong,
		Run:   runInit,
	}

	addCmd = &cobra.Command{
		Use:   "add [tag]",
		Short: T().AddShort,
		Long:  T().AddLong,
		Run:   runAdd,
	}

	pushCmd = &cobra.Command{
		Use:   "push [tag]",
		Short: T().PushShort,
		Long:  T().PushLong,
		Run:   runPush,
	}

	listCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   T().ListShort,
		Long:    T().ListLong,
		Run:     runList,
	}

	rmCmd = &cobra.Command{
		Use:   "rm [tag]",
		Short: T().RmShort,
		Long:  T().RmLong,
		Args:  cobra.ExactArgs(1),
		Run:   runRm,
	}

	// Initialize lang command
	langCmd = &cobra.Command{
		Use:   "lang [language]",
		Short: "Set or display current language",
		Long:  "Set the language to 'en' for English or 'zh' for Chinese, or display current language if no argument provided.",
		Args:  cobra.MaximumNArgs(1),
		Run:   runLang,
	}

	pushCmd.Flags().BoolVar(&pushAll, "all", false, T().PushAllFlag)

	rootCmd.AddCommand(initCmd, addCmd, pushCmd, listCmd, rmCmd, langCmd)
}

// Note: reinitializeCommands function removed to avoid circular dependency
// Language changes will take effect on next command execution

func runInit(cmd *cobra.Command, args []string) {
	tags, err := readTags()
	if err != nil {
		fmt.Printf(T().ErrorReadingRtagFile+"\n", err)
		return
	}

	if len(tags) == 0 {
		fmt.Println(T().RtagFileEmptyOrNotExist)
		interactiveAddTag()
	} else {
		fmt.Println(T().CurrentTags)
		for _, tag := range tags {
			fmt.Printf("  - %s\n", tag)
		}
	}
}

func runAdd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		// 交互式模式
		interactiveAddTag()
	} else {
		// 直接添加指定的 tag
		tag := args[0]
		if err := addTag(tag); err != nil {
			fmt.Printf(T().AddTagFailed+"\n", err)
		} else {
			fmt.Printf(T().AddTagSuccess+"\n", tag)
		}
	}
}

func runPush(cmd *cobra.Command, args []string) {
	if pushAll {
		// 推送所有 tags
		tags, err := readTags()
		if err != nil {
			fmt.Printf(T().ReadTagsFailed+"\n", err)
			return
		}

		if len(tags) == 0 {
			fmt.Println(T().NoTagsFound)
			return
		}

		pushTags(tags)
	} else if len(args) > 0 {
		// 推送指定的 tag
		tag := args[0]
		tags, err := readTags()
		if err != nil {
			fmt.Printf(T().ReadTagsFailed+"\n", err)
			return
		}

		// 检查 tag 是否存在
		found := false
		for _, t := range tags {
			if t == tag {
				found = true
				break
			}
		}

		if !found {
			fmt.Printf(T().TagNotExistInFile+"\n", tag)
			return
		}

		pushTags([]string{tag})
	} else {
		fmt.Println(T().SpecifyTagOrUseAll)
	}
}

func runList(cmd *cobra.Command, args []string) {
	tags, err := readTags()
	if err != nil {
		fmt.Printf(T().ReadTagsFailed+"\n", err)
		return
	}

	if len(tags) == 0 {
		fmt.Println(T().NoTagsFound)
		return
	}

	fmt.Println(T().AllTags)
	for _, tag := range tags {
		fmt.Printf("  - %s\n", tag)
	}
}

func runRm(cmd *cobra.Command, args []string) {
	tag := args[0]
	if err := removeTag(tag); err != nil {
		fmt.Printf(T().RemoveTagFailed+"\n", err)
	} else {
		fmt.Printf(T().RemoveTagSuccess+"\n", tag)
	}
}

func readTags() ([]string, error) {
	file, err := os.Open(rtagFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tags []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			tags = append(tags, line)
		}
	}

	return tags, scanner.Err()
}

func writeTags(tags []string) error {
	file, err := os.Create(rtagFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, tag := range tags {
		if _, err := fmt.Fprintln(file, tag); err != nil {
			return err
		}
	}

	return nil
}

func addTag(tag string) error {
	tags, err := readTags()
	if err != nil {
		return err
	}

	// 检查 tag 是否已存在
	for _, existingTag := range tags {
		if existingTag == tag {
			return fmt.Errorf(T().TagAlreadyExists, tag)
		}
	}

	tags = append(tags, tag)
	return writeTags(tags)
}

func removeTag(tag string) error {
	tags, err := readTags()
	if err != nil {
		return err
	}

	// 查找并删除 tag
	newTags := []string{}
	found := false
	for _, existingTag := range tags {
		if existingTag != tag {
			newTags = append(newTags, existingTag)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf(T().TagNotExist, tag)
	}

	return writeTags(newTags)
}

func interactiveAddTag() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(T().EnterTag)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf(T().ReadInputFailed+"\n", err)
			return
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println(T().Exit)
			return
		}

		if input == "" {
			fmt.Println(T().TagCannotBeEmpty)
			continue
		}

		if err := addTag(input); err != nil {
			fmt.Printf(T().AddTagFailed+"\n", err)
		} else {
			fmt.Printf(T().AddTagSuccess+"\n", input)
		}

		fmt.Print(T().ContinueAdding)
		continueInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf(T().ReadInputFailed+"\n", err)
			return
		}

		continueInput = strings.TrimSpace(strings.ToLower(continueInput))
		if continueInput != "y" && continueInput != "yes" {
			fmt.Println(T().Exit)
			return
		}
	}
}

func pushTags(tags []string) {
	currentTime := time.Now().Format("200601021504") // YYYYMMDDHHMM

	fmt.Printf(T().StartPushingTags+"\n", currentTime)

	for _, tag := range tags {
		gitTag := fmt.Sprintf("release-%s-%s", currentTime, tag)
		fmt.Printf(T().CreateGitTag+"\n", gitTag)

		// 执行 git tag 命令
		if err := executeCommand("git", "tag", gitTag); err != nil {
			fmt.Printf(T().CreateTagFailed+"\n", gitTag, err)
			continue
		}
	}

	// 推送所有 tags 到远程仓库
	fmt.Println(T().PushingTagsToRemote)
	if err := executeCommand("git", "push", "origin", "--tags"); err != nil {
		fmt.Printf(T().PushTagsFailed+"\n", err)
	} else {
		fmt.Println(T().PushTagsSuccess)
	}
}

func executeCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
