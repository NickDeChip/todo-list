import React, { Dispatch, SetStateAction, useEffect, useState } from 'react';
import { Btn } from './btn';
import { InputBox } from './input-box';

export function SignUp() {
  let [username, setUsername] = useState("")
  let [email, setEmail] = useState("")
  let [password, setPassword] = useState("")
  let [passwordCheck, setPasswordCheck] = useState("")
  let [err, setErr] = useState("")

  const CreateUser = () => {
    setErr("")

    if (password !== passwordCheck) {
      setErr("!Passwords Do Not Match!")
      return
    }
    (async () => {
      const res = await fetch("http://localhost:6969/signup", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Accept": "application/json"
        },
        credentials: 'include',
        body: JSON.stringify({
          username,
          email,
          password
        })
      })
      if (res.ok) {
        window.location.replace("/todo")
        return
      }
      const err = await res.text()
      setErr(err)
    })()
  }

  return <>
    <div className='flex justify-center mt-40 items-center'>
      <div className='w-7/12 lg:w-6/12 xl:w-5/12 border-[8px] border-black rounded-2xl h-auto font-seri md:text-xl xl:text-2xl'>
        {err != "" ?
          <div
            className='flex justify-center mt-3 text-2xl font-bold font-serif text-red-600 text-center'
          >
            {err}
          </div>
          : <></>}
        <InputBox
          vaule={username}
          label='Username -'
          placeholder='Write Your Username Here!'
          setVaule={setUsername}
          type='text'
        />
        <InputBox
          vaule={email}
          label='Email -'
          placeholder='Write Your Email Here!'
          setVaule={setEmail}
          type='text'
        />
        <InputBox
          vaule={password}
          label='Password -'
          placeholder='Write Your Password Here!'
          setVaule={setPassword}
          type='password'
        />
        <InputBox
          vaule={passwordCheck}
          label='Confirm Password -'
          placeholder='Write Your Password Here, AGAIN!'
          setVaule={setPasswordCheck}
          type='password'
        />
        <div className='flex justify-center'>
          <Btn
            label='LOGIN'
            onClickEvent={() => { window.location.replace("/signin") }}
            style='mr-4 p-1 mb-2'
          />
          <Btn
            label='SIGNUP'
            onClickEvent={CreateUser}
            style='ml-4 p-1 mb-2'
          />
        </div>
      </div>
    </div>
  </>
}
