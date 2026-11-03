package main

type MacOsComputer struct {
	Computer
}

func MakeMacOS() IComputer {
	return &MacOsComputer{
		Computer: Computer{
			os:    "MacOs",
			price: 3000,
		},
	}
}

type WindowsComputer struct {
	Computer
}

func MakeWindows() IComputer {
	return &WindowsComputer{
		Computer: Computer{
			os:    "Windows",
			price: 2000,
		},
	}
}

type LinuxComputer struct {
	Computer
}

func MakeLinux() IComputer {
	return &LinuxComputer{
		Computer: Computer{
			os:    "Linux",
			price: 1000,
		},
	}
}
