import React, { useEffect, useState } from 'react';
import { TodoList } from './components/todo-list';
import { TodoPrompt } from './components/todo-prompt';

export type TodoItem = {
  id: number,
  info: string,
}

function App() {
  let [todos, setTodos] = useState<TodoItem[]>([])

  useEffect(() => {
    (async () => {
      const res = await fetch("http://localhost:6969/todos");
      if (!res.ok) {
        console.error(res);
        return []
      }
      const todoItems = (await res.json()) as TodoItem[];
      setTodos(todoItems)
    })();
  }, [])

  return (
    <>
      <TodoPrompt todos={todos} setTodos={setTodos} />
      <TodoList todos={todos} setTodos={setTodos} />
    </>
  );
}

export default App;

