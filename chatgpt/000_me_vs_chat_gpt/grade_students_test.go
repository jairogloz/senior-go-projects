package chatgpt_test

import (
	chatgpt "github.com/JairoGLoz/senior-go-projects/senior-go-projects/chatgpt/000_me_vs_chat_gpt"
	"testing"
)

func BenchmarkGradingStudents1(b *testing.B) {

	a := []int32{73, 67, 38, 33}

	for i := 0; i < b.N; i++ {
		chatgpt.GradingStudents1(a)
	}
}

func BenchmarkGradingStudents2(b *testing.B) {

	a := []int32{73, 67, 38, 33}

	for i := 0; i < b.N; i++ {
		chatgpt.GradingStudents2(a)
	}
}
