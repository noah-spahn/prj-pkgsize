package pkgsize

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

// ComputePackageSizes takes a directory path as input and returns a map of package names to their sizes in lines of code.
func ComputePackageSizes(rootDir string) (map[string]int, error) {
	// Initialize a map to store the package sizes
	pkgSizes := make(map[string]int)

	// Walk the directory tree to find all Go files
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".go" {
			// Parse the Go file
			fset := token.NewFileSet()
			file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if err != nil {
				return err
			}

			// Compute the size of the package
			size := computePackageSize(fset, file)

			// Get the package name
			pkgName := file.Name.Name

			// Update the package size
			pkgSizes[pkgName] += size
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return pkgSizes, nil
}

func computePackageSize(fset *token.FileSet, file *ast.File) int {
	size := 0

	// Traverse the AST and count the lines of each function and variable declaration
	ast.Inspect(file, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.FuncDecl:
			funcNode := n.(*ast.FuncDecl)
			if funcNode.Body != nil {
				startLine := fset.Position(funcNode.Pos()).Line
				endLine := fset.Position(funcNode.End()).Line
				size += endLine - startLine + 1
			}
		case *ast.ValueSpec:
			valueNode := n.(*ast.ValueSpec)
			if valueNode.Values == nil {
				startLine := fset.Position(valueNode.Pos()).Line
				endLine := fset.Position(valueNode.End()).Line
				size += endLine - startLine + 1
			}
		}
		return true
	})

	return size
}
