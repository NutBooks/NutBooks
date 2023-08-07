import React from 'react';
import MenuItems from './MenuItems';
import Link from 'next/dist/client/link';
import {
  AiOutlineLogin,
  AiOutlineUserAdd,
  AiOutlineWhatsApp,
} from 'react-icons/ai';

function NavBar() {
  return (
    <nav className="grid grid-cols-3 items-center  px-[6rem]">
      <div>
        <Link href="/">
          <img src="/logo/Nutbooks_logo.png" className="w-[12rem]"></img>
        </Link>
      </div>
      <div className="grid grid-cols-3 font-PreB text-center">
        <Link href="/">
          <MenuItems href="/">
            <span className="text-[1.125rem]">넛북스</span>
          </MenuItems>
        </Link>
        <Link href="/insight">
          <MenuItems href="/insight">
            <span className="text-[1.125rem]">인사이트</span>
          </MenuItems>
        </Link>
        <Link href="/list">
          <MenuItems href="/list">
            <span className="text-[1.125rem]">모아보기</span>
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
        <Link href="/">
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
