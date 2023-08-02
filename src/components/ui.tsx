import { useEffect, useState } from "react";

type username = {
  username: string
}

export function UI() {
  let [name, setName] = useState<username>()

  useEffect(() => {
    (async () => {
      const res = await fetch("http://localhost:6969/user/username", { credentials: "include" });
      if (!res.ok) {
        console.error(res);
        return ""
      }
      const username = (await res.json()) as username;
      setName(username)
    })();
  }, [])

  return <>
    <div className="flex flex-row justify-between bg-gray-800 font-serif">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        strokeWidth="1.5"
        stroke="currentColor"
        className="w-6 h-6 mt-2 ml-3 text-white">
        <path
          strokeLinecap="round"
          strokeLinejoin="round"
          d="M15.75 6a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0zM4.501 20.118a7.5 7.5 0 0114.998 0A17.933 17.933 0 0112 21.75c-2.676 0-5.216-.584-7.499-1.632z"
        />
      </svg>
      <h1 className="mt-2 ml-3 text-white text-lg self-start">
        {name?.username}
      </h1>
      <button
        className='border-[3px] border-red-800 rounded-full bg-red-500 text-white transition ease-in-out duration-500 transform hover:scale-125 shadow-black my-1 mr-1 p-1 hover:bg-white hover:text-red-500'
        onClick={(_) => {
          document.cookie = "jwt=;expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
          window.location.reload()
        }}
      >
        SIGN OUT
      </button>
    </div>
  </>
}
