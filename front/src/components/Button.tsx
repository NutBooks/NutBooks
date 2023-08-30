import { backgroundColor, txtColor } from './color';

interface ButtonProps {
  innerValue: string;
  bgColor: string;
  textColor: string;
}

function RoundSm({ bgColor, textColor, innerValue }: ButtonProps) {
  return (
    <div
      className={` flex justify-center m-auto items-center text-center rounded-[20rem] px-3 w-[fit-content] h-[1.87rem] text-[1rem] bg-${bgColor} text-${textColor}`}
    >
      {innerValue}
    </div>
  );
}

export default RoundSm;
