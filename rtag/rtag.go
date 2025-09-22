package rtag

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const rtagFile = ".rtag"

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize .rtag file",
	Long:  `.rtag 文件用于存储已发布的 tag。`,
	Run:   runInit,
}

var addCmd = &cobra.Command{
	Use:   "add [tag]",
	Short: "Add a new tag",
	Long:  `Add a new tag to the .rtag file. If no tag is provided, interactive mode will be used.`,
	Run:   runAdd,
}

var pushCmd = &cobra.Command{
	Use:   "push [tag]",
	Short: "Push tags to remote repository",
	Long:  `Push tags to remote repository. Use --all flag to push all tags.`,
	Run:   runPush,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tags",
	Long:  `List all tags from the .rtag file.`,
	Run:   runList,
}

var rmCmd = &cobra.Command{
	Use:   "rm [tag]",
	Short: "Remove a tag",
	Long:  `Remove a tag from the .rtag file.`,
	Args:  cobra.ExactArgs(1),
	Run:   runRm,
}

var pushAll bool

func init() {
	pushCmd.Flags().BoolVar(&pushAll, "all", false, "Push all tags")

	rootCmd.AddCommand(initCmd, addCmd, pushCmd, listCmd, rmCmd)
}

func runInit(cmd *cobra.Command, args []string) {
	tags, err := readTags()
	if err != nil {
		fmt.Printf("错误读取 .rtag 文件: %v\n", err)
		return
	}

	if len(tags) == 0 {
		fmt.Println(".rtag 文件为空或不存在")
		interactiveAddTag()
	} else {
		fmt.Println("当前 tags:")
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
			fmt.Printf("添加 tag 失败: %v\n", err)
		} else {
			fmt.Printf("成功添加 tag: %s\n", tag)
		}
	}
}

func runPush(cmd *cobra.Command, args []string) {
	if pushAll {
		// 推送所有 tags
		tags, err := readTags()
		if err != nil {
			fmt.Printf("读取 tags 失败: %v\n", err)
			return
		}

		if len(tags) == 0 {
			fmt.Println("没有找到任何 tags")
			return
		}

		pushTags(tags)
	} else if len(args) > 0 {
		// 推送指定的 tag
		tag := args[0]
		tags, err := readTags()
		if err != nil {
			fmt.Printf("读取 tags 失败: %v\n", err)
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
			fmt.Printf("Tag '%s' 不存在于 .rtag 文件中\n", tag)
			return
		}

		pushTags([]string{tag})
	} else {
		fmt.Println("请指定要推送的 tag 或使用 --all 标志")
	}
}

func runList(cmd *cobra.Command, args []string) {
	tags, err := readTags()
	if err != nil {
		fmt.Printf("读取 tags 失败: %v\n", err)
		return
	}

	if len(tags) == 0 {
		fmt.Println("没有找到任何 tags")
		return
	}

	fmt.Println("所有 tags:")
	for _, tag := range tags {
		fmt.Printf("  - %s\n", tag)
	}
}

func runRm(cmd *cobra.Command, args []string) {
	tag := args[0]
	if err := removeTag(tag); err != nil {
		fmt.Printf("删除 tag 失败: %v\n", err)
	} else {
		fmt.Printf("成功删除 tag: %s\n", tag)
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
			return fmt.Errorf("tag '%s' 已存在", tag)
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
		return fmt.Errorf("tag '%s' 不存在", tag)
	}

	return writeTags(newTags)
}

func interactiveAddTag() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("请输入一个 tag (或输入 'exit' 退出): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("读取输入失败: %v\n", err)
			return
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println("退出")
			return
		}

		if input == "" {
			fmt.Println("Tag 不能为空，请重新输入")
			continue
		}

		if err := addTag(input); err != nil {
			fmt.Printf("添加 tag 失败: %v\n", err)
		} else {
			fmt.Printf("成功添加 tag: %s\n", input)
		}

		fmt.Print("是否继续添加? (y/n): ")
		continueInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("读取输入失败: %v\n", err)
			return
		}

		continueInput = strings.TrimSpace(strings.ToLower(continueInput))
		if continueInput != "y" && continueInput != "yes" {
			fmt.Println("退出")
			return
		}
	}
}

func pushTags(tags []string) {
	currentTime := time.Now().Format("200601021504") // YYYYMMDDHHMM

	fmt.Printf("开始推送 tags (时间戳: %s)...\n", currentTime)

	for _, tag := range tags {
		gitTag := fmt.Sprintf("release-%s-%s", currentTime, tag)
		fmt.Printf("创建 git tag: %s\n", gitTag)

		// 执行 git tag 命令
		if err := executeCommand("git", "tag", gitTag); err != nil {
			fmt.Printf("创建 tag %s 失败: %v\n", gitTag, err)
			continue
		}
	}

	// 推送所有 tags 到远程仓库
	fmt.Println("推送 tags 到远程仓库...")
	if err := executeCommand("git", "push", "origin", "--tags"); err != nil {
		fmt.Printf("推送 tags 失败: %v\n", err)
	} else {
		fmt.Println("成功推送所有 tags")
	}
}

func executeCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
