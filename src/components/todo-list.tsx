import React, { Dispatch, SetStateAction } from "react";
import { TodoItem } from "../App";

type todoProps = {
  todos: TodoItem[];
  setTodos: Dispatch<SetStateAction<TodoItem[]>>;
}

export function TodoList(props: todoProps) {

  const removeTodo = (id: number) => {
    (async () => {
      const res = await fetch("http://localhost:6969/todo", {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ id: id })
      })
      if (!res.ok) {
        // Show Error
        console.error(res)
      }
      props.setTodos(props.todos.filter((v) => v.id !== id))
    })()
  }

  return <>
    <div className="flex flex-col items-center">
      <h1
        className="underline font-serif text-6xl font-bold mt-6 cursor-default md:text-8xl">
        YOUR TODOS
      </h1>
      <h1
        className="font-serif text-6xl font-bold mb-6 cursor-default md:text-8xl bg-gradient-to-t from-black to-transparent text-transparent bg-clip-text scale-y-[-1]">
        YOUR TODOS
      </h1>
    </div>
    <div className="flex flex-col items-center lg:grid lg:grid-cols-3">{
      props.todos.map((v, i, _) =>
        <div
          className="todo-box"
          key={i}
          onClick={(_) => { removeTodo(v.id) }}
        >
          <span className="m-3">{v.info}</span>
        </div>
      )
    }</div>
  </>
}

