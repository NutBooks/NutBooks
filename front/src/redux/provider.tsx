// store컴포넌트가 client-side에 있어야 함을 표시
"use client";

import { store } from "./store";
import { Provider } from "react-redux";

export function Providers({ children }: { children: React.ReactNode }) {
  return <Provider store={store}>{children}</Provider>;
}