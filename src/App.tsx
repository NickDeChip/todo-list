import React from 'react';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { ErrorHandle } from './components/root-error';
import { SignIn } from './components/signin';
import { SignUp } from './components/signup';
import { Todo } from './components/todo';

function App() {
  return <RouterProvider router={router} />
}

const router = createBrowserRouter([
  {
    errorElement: <ErrorHandle />
  },
  {
    path: "/todo",
    element: <Todo />
  },
  {
    path: "/signup",
    element: <SignUp />
  },
  {
    path: "/signin",
    element: <SignIn />
  }
])

export default App;

