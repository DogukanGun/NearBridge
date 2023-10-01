import { SessionProvider } from "next-auth/react";

export default function ProviderWrapper({
  children,
  session
}: {
  children: React.ReactNode;
  session: any;
}): JSX.Element {
  return (
    <SessionProvider session={session}>
      {children}
    </SessionProvider>
  );
}
