import { Providers } from '../redux/provider';
import '@/styles/globals.css';

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>
        <header>
          <nav></nav>
        </header>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
