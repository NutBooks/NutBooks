import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { posts } from '../dummyData'; // 더미 데이터를 정의한 파일의 경로를 넣어주세요

export const postApi = createApi({
  reducerPath: 'postApi',
  refetchOnFocus: true,
  baseQuery: fetchBaseQuery({
    baseUrl: 'https://example.com/', // 실제 API 주소를 여기에 넣어주세요
  }),
  endpoints: (builder) => ({
    getPosts: builder.query<Post[], null>({
      query: () => 'posts', // 실제 API 엔드포인트 주소를 여기에 넣어주세요
    }),
    // 다른 엔드포인트 정의도 가능
  }),
});
export const { useGetPostsQuery } = postApi;
