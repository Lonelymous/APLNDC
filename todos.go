package main

var DB []TodoItem = make([]TodoItem, 0)
var PrimaryKey int = 1

func GetTodoItems() ([]TodoItem, error) {
	result := make([]TodoItem, 0)

	// SELECT * FROM TODOS

	for _, item := range DB {
		result = append(result, item)
	}

	return result, nil
}

func GetTodoItem(id int) (*TodoEditor, error) {
	result := &TodoEditor{}

	// SELECT * FROM TODOS WHERE TodoId = :Id

	for _, item := range DB {
		if item.TodoId == id {
			result.Done = item.Done
			result.Text = item.Text
			break
		}
	}

	return result, nil
}

func CreateTodoItem(editor TodoEditor) error {
	// INSERT INTO TODOS(`Text`,`Done`) VALUES(:text, :done)
	DB = append(DB, TodoItem{
		TodoId: PrimaryKey,
		Text:   editor.Text,
		Done:   editor.Done,
	})
	PrimaryKey++
	return nil
}

func UpdateTodoItem(editor TodoEditor, id int) error {
	// UPDATE TODOS SET `Text`=:text, `Done`=:done WHERE `TodoId`=:id
	for i := range DB {
		if DB[i].TodoId == id {
			DB[i].Done = editor.Done
			DB[i].Text = editor.Text
			break
		}
	}
	return nil
}

func DeleteTodoItem(id int) error {
	// DELETE FROM TODOS WHERE `TodoId`=:id
	for i := range DB {
		if DB[i].TodoId == id {
			DB = append(DB[:i], DB[i+1:]...)
			break
		}
	}
	return nil
}
