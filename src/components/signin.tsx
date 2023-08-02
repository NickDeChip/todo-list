import React, { useState } from 'react';
import { Btn } from './btn';
import { InputBox } from './input-box';


export function SignIn() {
  let [email, setEmail] = useState("")
  let [password, setPassword] = useState("")
  let [err, setErr] = useState("")

  const LogInUser = () => {
    setErr("");
    (async () => {
      const res = await fetch("http://localhost:6969/signin", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Accept": "application/json"
        },
        credentials: 'include',
        body: JSON.stringify({
          email,
          password
        })
      })
      if (res.ok) {
        console.log(res.text)
        window.location.replace("/todo")
        return
      }
      const err = await res.text();
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
        <div className='flex justify-center'>
          <Btn
            label='SIGNUP'
            onClickEvent={() => { window.location.replace("/signup") }}
            style='mr-4 p-1 mb-2'
          />
          <Btn
            label='LOGIN'
            onClickEvent={LogInUser}
            style='ml-4 p-1 mb-2'
          />
        </div>
      </div>
    </div>
  </>
}
