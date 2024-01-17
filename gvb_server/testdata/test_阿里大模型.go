package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type RequestBody struct {
	Model string `json:"model"`
	Input struct {
		Messages []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"messages"`
	} `json:"input"`
	Parameters struct {
	} `json:"parameters"`
}

type Output struct {
	FinishReason string `json:"finish_reason"`
	Text         string `json:"text"`
}

type Usage struct {
	TotalTokens  int `json:"total_tokens"`
	OutputTokens int `json:"output_tokens"`
	InputTokens  int `json:"input_tokens"`
}

type Response struct {
	Output Output `json:"output"`
	Usage  Usage  `json:"usage"`
}

func getChat(content string) (string, error) {
	// 创建一个通道用于接收信号
	done := make(chan bool)
	// 启动B函数，并传递通道作为参数
	go func() {
		spinner(done)
	}()
	// getChat函数结束时发送信号给通道
	defer func() {
		done <- true
	}()

	var body RequestBody
	body.Model = "qwen-max"
	body.Input.Messages = append(body.Input.Messages, struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}{Role: "system", Content: "You are a helpful assistant."})

	body.Input.Messages = append(body.Input.Messages, struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}{Role: "user", Content: content})

	jsonData, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	req, err := http.NewRequest("POST", "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "sk-ef191dc6db5e497f973a9e616f41120f")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	var response Response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return response.Output.Text, nil
}

func spinner(done chan bool) {
	// B函数的死循环逻辑
	for {
		select {
		case <-done: // 接收到信号，退出循环
			return
		default:
			// 执行B函数的逻辑
			for _, r := range `-\|/` {
				fmt.Printf("\r思考中%c", r)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

func main() {
	for {
		fmt.Print("请输入你的问题：")
		reader := bufio.NewReader(os.Stdin)
		content, _ := reader.ReadString('\n')
		content = strings.TrimSpace(content)
		if content == "退出" || content == "exit" {
			break
		}

		result, err := getChat(content)
		if err != nil {
			fmt.Println("\r调用失败！")
			continue
		}
		fmt.Println("\rAI大模型：", result)
		fmt.Println()
	}
}
