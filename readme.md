# Letter Avatar Generator

## Introduction
The Letter Avatar Generator is a Go application that creates personalized avatars using the first letter of a user's name. Users can choose a background color and avatar shape before generating their avatar. This project makes use of the `letteravatarapi` custom package for generating an image with a letter set against a customizable background. It's perfect for creating unique avatars for profiles, forums, or any platform that supports personalized images.

## Features
- Allows users to choose a background color (Magenta, Green, Blue).
- Users can select an avatar shape (Circle or Square).
- Generates avatars with customizable size and color settings.
- Saves the avatar in PNG format.

## Requirements
- Go (version 1.15 or newer is recommended)
- The `gg` package for Go graphics

## Installation

Ensure Go is installed on your system by downloading it from [the official Go website](https://golang.org/dl/).

Clone the repository to get started with the Letter Avatar Generator:

```bash
git clone https://github.com/modevsty/go-letteravatar.git
cd go-letteravatar
```
Install the required Go package:

```bash
go get github.com/fogleman/gg
```
## Usage
To run the application, navigate to the project directory and build the application:

```bash
go build
```
Then run the following command:
```bash
./letteravatar
```
The application will guide you through the following steps:

1. Enter your first and last name.
2. Pick a background color from the provided list by entering the corresponding number.
3. Choose a shape for your avatar by selecting from the options provided.
4. After these steps, the program generates an avatar based on your selections and saves it as a PNG file in the current directory.

## Customizing Avatar
The letteravatarapi allows for customization of the avatar's shape, size, and colors. Here are some methods you can use:

- WithShape("circle" or "rectangle"): Defines the shape of the avatar.
- WithSize(int): Sets the size of the avatar.
- SetColor(background, foreground string): Sets the background and foreground colors.

These methods can be chained together when creating a new LetterAvatar instance.

## Contributing
We welcome contributions to the Letter Avatar Generator project. Feel free to fork the repository, make your changes, and submit a pull request. For any issues or suggestions, please open an issue in the GitHub repository.