package main


func createFileList(files []os.FileInfo) []FileEntry {
	size := len(files)
	entryList := make([]FileEntry, size)
	for i, v := range files {
		entryList[i] = FileEntry{v.Name(), v.Size()}
	}
	return entryList
}
