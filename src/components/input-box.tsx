import { Dispatch, HTMLInputTypeAttribute, SetStateAction } from "react"

type Props = {
  vaule: string,
  label: string,
  placeholder: string,
  setVaule: Dispatch<SetStateAction<string>>
  type: HTMLInputTypeAttribute
}

export function InputBox(props: Props) {
  return <>
    <div className='flex flex-col items-center h-1/5 bg-sky-500 border-2 rounded-lg border-black m-5'>
      <label
        className='h-1/3 mt-2 self-start ml-10  text-blue-200'>
        {props.label}
      </label>
      <input
        className='h-2/3 w-4/6 sm:w-5/6 mb-3 align-middle focus:border-sky-200 rounded-lg border-[3px] border-zinc-300 transition ease-in-out duration-500 transform hover:scale-110 hover:shadow-black hover:shadow shadow-lg p-1'
        placeholder={props.placeholder}
        type={props.type}
        onChange={(e) => props.setVaule(e.target.value)}
        value={props.vaule}
      />
    </div>
  </>
}
