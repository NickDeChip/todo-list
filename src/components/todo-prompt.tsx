import React, { Dispatch, SetStateAction, useState } from 'react';
import { TodoItem } from './todo';

type todoProps = {
  todos: TodoItem[];
  setTodos: Dispatch<SetStateAction<TodoItem[]>>
}

type ID = {
  id: number
}

export function TodoPrompt(props: todoProps) {
  let [data, setData] = useState("");

  const saveTodo = (v: string) => {
    (async () => {
      const res = await fetch("http://localhost:6969/api/todo", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ info: v }),
        credentials: 'include',
      })
      const id = (await res.json()) as ID;
      const todo: TodoItem = {
        id: id.id,
        info: v
      }
      props.setTodos([...props.todos, todo])
    })()
  }

  return <>
    <div className='flex flex-col items-center justify-center'>
      <div className='rainbow mt-4 h-24 border-2 border-black rounded-md shadow-sm shadow-black'>
        <label
          className='text-3xl font-serif ml-2 cursor-defaul font-bold'>
          To Do:
        </label>
        <input
          className='mx-4 mb-2 text-2xl placeholder:italic placeholder:text-slate-400 block rounded-md hover:shadow-black hover:shadow-md focus:border-sky-500 transition ease-in-out duration-500 transform hover:scale-110 p-1 border-2 border-black hover:placeholder:text-sky-400'
          placeholder='Add Your Next Todo'
          type="text"
          value={data}
          onChange={(e) => setData(e.target.value)}
          onKeyUp={
            (e) => {
              e.key === 'Enter' && saveTodo(data)
              e.key === 'Enter' && setData("");
            }}
        />
      </div>
    </div>
  </>
}
