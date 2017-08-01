package cmd

type Conf struct {
	StorageName   string
	ResourceGroup string
	Location      string
	ShareName     string
}

const Config = `twitter: {{.Twitter}}
author: {{.Author}}
email: {{.Email}}
moduledir: {{.ModuleDir}}
coursedir: {{.CourseDir}}
`
