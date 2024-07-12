package constants

type NavItemsType []struct {
	name string
	link string
}

var NavItems = NavItemsType{
	{
		name: "About",
		link: "/about",
	},
	{
		name: "Portfolio",
		link: "/portfolio",
	},
	{
		name: "Others",
		link: "/about#others",
	},
	{
		name: "Contact",
		link: "/#contact",
	},
}
