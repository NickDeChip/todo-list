import { useEffect, useState } from "react";
import { OfferLogin } from "./offer-login";
import { TodoList } from "./todo-list";
import { TodoPrompt } from "./todo-prompt";
import { UI } from "./ui";

export type TodoItem = {
  id: number,
  info: string,
}

export function Todo() {
  let [todos, setTodos] = useState<TodoItem[]>([])
  let [logedIn, setLogedIn] = useState(false)

  useEffect(() => {
    (async () => {
      const res = await fetch("http://localhost:6969/api/todo/all", { credentials: "include" });
      if (!res.ok) {
        if (res.status === 401) { setLogedIn(false) };
        return []
      }
      const todoItems = (await res.json()) as TodoItem[];
      setLogedIn(true)
      setTodos(todoItems)
    })();
  }, [])

  return (
    <>
      {logedIn ? <>
        <UI />
        <TodoPrompt todos={todos} setTodos={setTodos} />
        <TodoList todos={todos} setTodos={setTodos} />
      </>
        : <OfferLogin />
      }
    </>
  );
}
