import { Providers } from '../redux/provider';
import '@/styles/globals.css';
import NavBar from '@/components/Navbar';

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>
        <header>
          <nav>
            <NavBar />
          </nav>
        </header>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
