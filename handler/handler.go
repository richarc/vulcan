package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/richarc/vulcan/model"
)

const claudePromptFormat = "\n\nHuman: %s\n\nAssistant:"

type Prompt struct {
	Question string `form:"question"`
}

func Hello(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{}, "layouts/main")
}

func Chat(c *fiber.Ctx) error {
	return c.Render("chat", fiber.Map{}, "layouts/main")
}

func Send(c *fiber.Ctx) error {
	//retrieve the question from our form
	prompt := new(Prompt)
	if err := c.BodyParser(prompt); err != nil {
		return err
	}

	//call a model with the chat input and get a sttring response
	msg := fmt.Sprintf(claudePromptFormat, prompt.Question)

	ans, err := model.AskLLM(msg)
	if err != nil {
		log.Fatal("Could not get response from LLM", err)
	}

	//return the template populating the response
	return c.Render("send", fiber.Map{"Answer": ans, "Prompt": prompt.Question}, "layouts/main")
}
