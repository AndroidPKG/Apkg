package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	// Get the current user
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user:", err)
		return
	}

	// Check if user is "shell"
	if currentUser.Username == "shell" {
		fmt.Println("! Shell user detected")
		setupShellUser()
	}

	// Check if user is "u0_*" (Termux user)
	if filepath.HasPrefix(currentUser.Username, "u0_") {
		fmt.Println("! Termux user detected")
		setupTermuxUser()
	}
}

// Function for error checking and directory creation
func createDir(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory %s: %v\n", path, err)
	} else {
		fmt.Printf("> Created directory: %s\n", path)
	}
}

// Setting up directories and variables for the "shell" user
func setupShellUser() {
	fmt.Println("> Creating essential directories...")
	createDir("/storage/emulated/0/Apkg/apk")
	createDir("/storage/emulated/0/Apkg/gzip")
	createDir("/data/local/tmp/apkg/bin")
	createDir("/data/local/tmp/apkg/lib")

	fmt.Println("> Exporting Variables...")
	os.Setenv("Storage", "/storage/emulated/0")
	os.Setenv("HomeDir", "/data/local/tmp")
	os.Setenv("ApkgStore", "/storage/emulated/0/Apkg")
	os.Setenv("ApkgHome", "/data/local/tmp/apkg")
	os.Setenv("apk", "/storage/emulated/0/Apkg/apk")
	os.Setenv("gzip", "/storage/emulated/0/Apkg/gzip")
	os.Setenv("bin", "/data/local/tmp/apkg/bin")
	os.Setenv("lib", "/data/local/tmp/apkg/lib")
}

// Setting up directories and variables for the "u0_*" (Termux) user
func setupTermuxUser() {
	fmt.Println("> Creating essential directories...")
	createDir("/storage/emulated/0/Apkg/apk")
	createDir("/storage/emulated/0/Apkg/gzip")
	createDir("/data/data/com.termux/files/home/Apkg/bin")
	createDir("/data/data/com.termux/files/home/Apkg/lib")

	fmt.Println("> Exporting Variables...")
	os.Setenv("Storage", "/storage/emulated/0")
	os.Setenv("HomeDir", "/data/data/com.termux/files/home")
	os.Setenv("ApkgStore", "/storage/emulated/0/Apkg")
	os.Setenv("ApkgHome", "/data/data/com.termux/files/home/Apkg")
	os.Setenv("apk", "/storage/emulated/0/Apkg/apk")
	os.Setenv("gzip", "/storage/emulated/0/Apkg/gzip")
	os.Setenv("bin", "/data/data/com.termux/files/home/Apkg/bin")
	os.Setenv("lib", "/data/data/com.termux/files/home/Apkg/lib")

	// Append bin directory to PATH
	path := os.Getenv("PATH")
	os.Setenv("PATH", fmt.Sprintf("%s:/data/data/com.termux/files/home/Apkg/bin", path))
}
