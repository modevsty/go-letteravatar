# Letter Avatar Generator

## Introduction
The Letter Avatar Generator is a Go application designed to create personalized avatars using the first letter of a user's name. Users have the flexibility to choose background colors, avatar shapes, size and the file format for saving the generated avatar. This project utilizes the `letteravatarapi` custom package for generating images with a letter set against a customizable background, making it ideal for creating unique avatars for various online platforms.

## Features
- Users can choose a background color from predefined options.
- Shape selection allows for either a circle or square avatar.
- Size customization enables specifying the avatar dimension within a range.
- Supports saving the avatar in either PNG or JPG formats.

## Requirements
- Go (version 1.15 or newer is recommended)
- The `gg` package for Go graphics

## Installation

Ensure you have Go installed on your system by downloading it from [the official Go website](https://golang.org/dl/).

Clone the repository to begin using the Letter Avatar Generator:

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
The application will guide you through a series of prompts:

1. Enter Your Name: Input your first and last name only.
2. Pick a Background Color: Choose from the list by entering the corresponding number.
3. Pick a Shape: Select your preferred shape for the avatar.
4. Pick a Size: Specify the avatar's size (in pixels), which must be between 16 and 1024.
4. Pick a Save Format: Choose between PNG or JPG formats for your avatar.
5. After these selections, the program generates your avatar according to your specifications and saves it in the current directory.



## Customizing Avatar
The letteravatarapi allows for customization of the avatar's shape, size, and colors. Here are some methods you can use:

- WithShape("circle" or "rectangle"): Defines the shape of the avatar.
- WithSize(int): Sets the size of the avatar.
- WithColor(bgColorHex string): Sets the background color.

These methods can be chained together when creating a new LetterAvatar instance.

## Contributing
We welcome contributions to the Letter Avatar Generator project. Feel free to fork the repository, make your changes, and submit a pull request. For any issues or suggestions, please open an issue in the GitHub repository.