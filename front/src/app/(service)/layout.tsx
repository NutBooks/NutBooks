import AddBookmarkBar from '@/components/AddBookmarkBar';

export default function ServiceLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <>
      <AddBookmarkBar />
      {children}
    </>
  );
}
