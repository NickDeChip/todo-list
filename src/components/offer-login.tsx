import { Btn } from "./btn";

export function OfferLogin() {
  return (
    <>
      <div className='flex justify-center mt-40 items-center'>
        <div className='w-7/12 lg:w-6/12 xl:w-5/12 border-[8px] border-black rounded-2xl h-auto font-seri md:text-xl xl:text-2xl text-center font-serif'>
          <h1 className="text-4xl">Looks like your not logged in!</h1>
          <p className="text-2xl mb-2">Try logging in you silly billy, or signing up if you don't all ready have a account</p>
          <div className='flex justify-center'>
            <Btn
              label='SIGNUP'
              onClickEvent={() => { window.location.replace("/signup") }}
              style='mr-4 p-1 mb-2'
            />
            <Btn
              label='LOGIN'
              onClickEvent={() => { window.location.replace("/signin") }}
              style='ml-4 p-1 mb-2'
            />
          </div>
        </div>
      </div>
    </>)
}
