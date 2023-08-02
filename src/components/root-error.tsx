import { isRouteErrorResponse, useRouteError } from "react-router-dom"

export function ErrorHandle() {
  let err = useRouteError()
  return (
    <div className="flex flex-col items-center justify-center text-center mt-40">
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" className="w-20 h-20 text-red-300 animate-bounce">
        <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126zM12 15.75h.007v.008H12v-.008z" />
      </svg>
      <h1 className="text-4xl">Dang! Looks like there was a Error</h1>
      <p className="text-slate-700 text-2xl">
        {
          isRouteErrorResponse(err)
            ? (err.error?.message || err.status)
            : "Unknown error message"
        }
      </p>
    </div>
  )
}
