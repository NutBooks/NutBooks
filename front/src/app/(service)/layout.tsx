import AddBookmarkBar from '@/components/clientComponents/AddBookmarkBar';

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
