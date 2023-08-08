'use client';

import React, { Children } from 'react';
import { usePathname } from 'next/navigation';

function MenuItems({ children, href }: { children: any; href?: string }) {
  const pathname = usePathname();
  const isActive = pathname === href;

  return (
    <div className={`${isActive ? 'text-grassgreen' : 'text-black'} `}>
      {children}
    </div>
  );
}

export default MenuItems;
