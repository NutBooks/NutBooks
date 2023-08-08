import React from 'react';
import MenuItems from './MenuItems';
import Link from 'next/dist/client/link';
import { INSIGHT, MAIN, LIST, ASK } from '@constants/urlpath';
import {
  MAIN_CONTENTS,
  SECOND_CONTENTS,
  THIRD_CONTENTS,
} from '@constants/rootLayout';
import {
  AiOutlineLogin,
  AiOutlineUserAdd,
  AiOutlineWhatsApp,
} from 'react-icons/ai';

function NavBar() {
  return (
    <nav className="grid grid-cols-3 items-center  px-[6rem]">
      <div>
        <Link href={MAIN}>
          <img src="/logo/Nutbooks_logo.png" className="w-[12rem]"></img>
        </Link>
      </div>
      <div className="grid grid-cols-3 font-PreB text-center">
        <Link href={MAIN}>
          <MenuItems href={MAIN}>
            <span className="text-[1.125rem]">{MAIN_CONTENTS}</span>
          </MenuItems>
        </Link>
        <Link href={INSIGHT}>
          <MenuItems href={INSIGHT}>
            <span className="text-[1.125rem]">{SECOND_CONTENTS}</span>
          </MenuItems>
        </Link>
        <Link href={LIST}>
          <MenuItems href={LIST}>
            <span className="text-[1.125rem]">{THIRD_CONTENTS}</span>
          </MenuItems>
        </Link>
      </div>
      <div className="flex justify-end">
        <Link href="/login" className="mr-6">
          <AiOutlineLogin className="w-full" size={20} />
          <span className="text-gray-600 text-[.95rem]">로그인</span>
        </Link>
        <Link href="/" className="mr-6">
          <AiOutlineUserAdd className="w-full " size={20} />
          <span className="text-gray-600 text-[.95rem]">회원가입</span>
        </Link>
        <Link href={ASK}>
          <div className="items-center text-center">
            <AiOutlineWhatsApp className="w-full" size={20} />
            <span className="text-gray-600 text-[.95rem]">문의하기</span>
          </div>
        </Link>
      </div>
    </nav>
  );
}

export default NavBar;
