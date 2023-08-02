import { MouseEventHandler } from "react"

type Props = {
  label: string,
  onClickEvent: MouseEventHandler,
  style: string,
}

export function Btn(props: Props) {
  return <>
    <div className='flex justify-center h-1/5'>
      <button
        className={'border-[3px] border-blue-500 rounded-full bg-sky-400 text-white transition ease-in-out duration-500 transform hover:scale-125 shadow-black hover:text-sky-400 hover:bg-white hover:shadow-black hover:shadow-lg ' + props.style}
        onClick={props.onClickEvent}
      >
        {props.label}
      </button>
    </div>
  </>
}
