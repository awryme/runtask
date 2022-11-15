package taskfile

import (
	"github.com/matryer/is"
	"testing"
)

func createTestTaskfile() *Taskfile {
	file := &Taskfile{}

	file.Set("test1", "command test1")
	file.Set("test2", "command test2")
	file.Set("test3", "command test3")

	return file
}

func TestTaskfile_Set(t *testing.T) {
	assert := is.New(t)

	file := createTestTaskfile()

	assert.Equal(len(file.Commands), 3)          // should have 3 commands
	assert.Equal(file.Commands[0].name, "test1") // 1 should be test1
	assert.Equal(file.Commands[1].name, "test2") // 2 should be test2
	assert.Equal(file.Commands[2].name, "test3") // 3 should be test3
}

func TestTaskfile_SetDefault(t *testing.T) {
	assert := is.New(t)

	file := createTestTaskfile()

	assert.True(file.SetDefault("test2"))        // should be able to set test2 as default
	assert.True(!file.SetDefault("nonexistent")) // should not be able to set "nonexistent"as default

	assert.Equal(len(file.Commands), 3)          // should have 3 commands
	assert.Equal(file.Commands[0].name, "test2") // 1 should be test2
	assert.Equal(file.Commands[1].name, "test1") // 2 should be test1
	assert.Equal(file.Commands[2].name, "test3") // 3 should be test3
}

func TestTaskfile_Delete(t *testing.T) {
	assert := is.New(t)

	file := createTestTaskfile()

	assert.True(file.Delete("test2"))        // should be able to delete test2
	assert.True(!file.Delete("nonexistent")) // should not be able to delete nonexistent

	assert.Equal(len(file.Commands), 2)          // should have 2 commands
	assert.Equal(file.Commands[0].name, "test1") // 1 should be test1
	assert.Equal(file.Commands[1].name, "test3") // 2 should be test3
}

func TestTaskfile_Get(t *testing.T) {
	assert := is.New(t)

	file := createTestTaskfile()

	cmd, ok := file.Get("test2")

	assert.True(ok)                                 // should have test2
	assert.Equal(cmd.Name(), "test2")               // name should test2
	assert.Equal(cmd.RawCommand(), "command test2") // command should be "command test2"
}

func TestTaskfile_DefaultInit(t *testing.T) {
	assert := is.New(t)

	file := createTestTaskfile()

	cmd, ok := file.Default()

	assert.True(ok)                                 // should have default
	assert.Equal(cmd.Name(), "test1")               // default name should test1
	assert.Equal(cmd.RawCommand(), "command test1") // command should be "command test1"
}

func TestTaskfile_DefaultWithSet(t *testing.T) {
	assert := is.New(t)

	file := createTestTaskfile()

	file.SetDefault("test3")

	cmd, ok := file.Default()

	assert.True(ok)                                 // should have default
	assert.Equal(cmd.Name(), "test3")               // default name should test3
	assert.Equal(cmd.RawCommand(), "command test3") // command should be "command test3"
}
