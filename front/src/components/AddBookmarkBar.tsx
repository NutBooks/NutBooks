'use client';

import React, { useState } from 'react';
import { AddBookmark } from '@utils/api';
import { ADD_BOOKMARK } from '@constants/serviceLayout';
import { BsPlusCircleFill } from 'react-icons/bs';

function AddBookmarkBar() {
  const [bookmarkUrl, setBookmarkUrl] = useState('');
  const [bookmarkTitle, setBookmarkTitle] = useState('');

  const submitHandler = () => {
    AddBookmark({ title: bookmarkTitle, link: bookmarkUrl });
  };

  return (
    <div>
      <form onSubmit={submitHandler}>
        <input type="text" placeholder={ADD_BOOKMARK.PLACEHOLDER} />
        <div className="w-10">
          <BsPlusCircleFill />
        </div>
      </form>
    </div>
  );
}

export default AddBookmarkBar;
