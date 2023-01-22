package standartstegano

import (
	"testing"
)

func TestPerformStandartEncryption(t *testing.T) {
	container := []byte("Hello, World! This is my first day here. I am so glad to see all of you")
	err := PerformStandartHiding(container, []byte("Love"))
	containerWithHiddenMessage := []byte("Hdlln, Vnrld !Thhs hs!lx girsu!e`y!idse. I!`l!ro glad to see all of you")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if string(container) != string(containerWithHiddenMessage) {
		t.Errorf("expected: %s\ngot: %s", containerWithHiddenMessage, container)
	}
}

func TestPerformStandarDecryption(t *testing.T) {
	container := []byte("Hdlln, Vnrld !Thhs hs!lx girsu!e`y!idse. I!`l!ro glad to see all of you")
	decryptedMessage, err := PerformStandartExtraction(container)
	message := "Love"
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if string(decryptedMessage) != message {
		t.Errorf("expected: %s\ngot: %s", decryptedMessage, message)
	}
}
