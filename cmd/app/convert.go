package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jimma/blogmarkdownconverter/pkg/blogindex"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:              "run <blog url>",
	Short:            "Convert blog entry to markdown file",
	Long:             `Run converter to convert blog entry to Markdown file`,
	Args:             cobra.MinimumNArgs(1),
	TraverseChildren: true,
	Run:              Convert,
}

func Convert(cmd *cobra.Command, args []string) {

	if err := doConvert(cmd, args); err != nil {
		handleErr(err)
	}

}

func doConvert(cmd *cobra.Command, args []string) error {
	outputdir, _ := cmd.Flags().GetString("output")
	if outputdir == "" {
		outputdir = "./"
	}
	absdir, absErr := filepath.Abs(outputdir)
	if absErr != nil {
		handleErr(absErr)
	}
	finfo, ferr := os.Stat(absdir)
	if os.IsNotExist(ferr) {
		return errors.New("Provided output dir: " + outputdir + " doesn't exist")
	}
	if !finfo.IsDir() {
		return errors.New("Provided output dir: " + outputdir + " is not a directory")

	}

	var index = blogindex.BlogIndex{
		URL:   args[0],
		XPath: blogindex.BlogIndexPath,
	}
	_, err := index.Parse()
	if err != nil {
		return errors.New(args[0] + " is not a valid blog url")
	}

	blogs := index.GetBlogEntry()
	for _, blog := range blogs {
		f, err := os.Create(absdir + "/" + blog.MDFileName + ".md")
		if err == nil {
			f.WriteString(blog.Content)
			f.Close()
		} else {
			return err
		}
	}
	fmt.Printf("%d blog entries are converted\n", len(blogs))
	return nil
}

func handleErr(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(-1)
}
