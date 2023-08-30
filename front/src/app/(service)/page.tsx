'use client';

import { useGetUsersQuery } from '@/redux/services/userApi';
import { useAppDispatch, useAppSelector } from '@/redux/hooks';
import { posts } from '@/redux/dummyData';
import RoundSm from '@/components/Button';

export default function Home() {
  const count = useAppSelector((state) => state.counterReducer.value);
  const dispatch = useAppDispatch();

  const { isLoading, isFetching, data, error } = useGetUsersQuery(null);

  return (
    <main style={{ maxWidth: 1200, marginInline: 'auto', padding: 20 }}>
      <div style={{ marginBottom: '4rem', textAlign: 'center' }}></div>
      {/* 하위 컴포넌트 */}
      <div
        style={{
          boxShadow: '0px 0px 15px lightgrey',
          width: '100%',
          padding: '3rem',
          borderRadius: '2rem',
        }}
      >
        <p className="font-PreB text-[1.4rem] mb-6">전체 넛북스</p>
        <table style={{ width: '100%' }} className="my-4">
          <thead className="bg-[#F8F8F8] h-[3.2rem] ">
            <tr className="text-[1.2rem] text-center font-PreR border-t-[1px] border-b-[1px] border-slate-500">
              <th>제목</th>
              <th>링크</th>
              <th>키워드</th>
              <th>요약상태</th>
            </tr>
          </thead>
          <tbody>
            {posts.map((post) => (
              <tr
                className="border-b-[1px]  border-t-[1px]  border-slate-300"
                key={post.id}
              >
                <td
                  style={{
                    margin: '1rem',
                    display: 'block',
                    width: '14rem',
                    textOverflow: 'ellipsis',
                    overflow: 'hidden',
                    whiteSpace: 'nowrap',
                  }}
                >
                  {post.title}
                </td>
                <td>
                  <a
                    style={{
                      display: 'block',
                      width: '17rem',
                      textOverflow: 'ellipsis',
                      overflow: 'hidden',
                      whiteSpace: 'nowrap',
                    }}
                    href={post.link}
                  >
                    {post.link}
                  </a>
                </td>
                <td className="flex">
                  {post.keywords.map((keyword, index) => (
                    <RoundSm
                      key={index}
                      textColor="[#FFFFFF]"
                      bgColor="[#D9D9D9]"
                      innerValue={keyword}
                    />
                  ))}
                </td>
                <td className="w-[10rem] ">
                  {post.status ? (
                    <RoundSm
                      textColor="[#75A86]"
                      bgColor="[#DAFEBC]"
                      innerValue="완료"
                    />
                  ) : (
                    <RoundSm
                      textColor="[#9BA5B7]"
                      bgColor="[#EEF1F4]"
                      innerValue="대기"
                    />
                  )}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </main>
  );
}
