"use client"
import CustomNavbar from '@/components/navbar/CustomNavbar'
import '../styles/globals.css'
import CustomFooter from '@/components/footer/CustomFooter'
import Provider from "@/app/context/client-provider"
import { getServerSession } from "next-auth/next"
import { authOptions } from "@/app/api/auth/[...nextauth]/route"
import { Montserrat } from 'next/font/google'
import { useState, useEffect } from 'react'


const montserrat = Montserrat({ subsets: ['latin'] })

export default function RootLayout({
  children
}: {
  children: React.ReactNode
}): React.ReactNode {
  const [session, setSession] = useState<any>(null);

  useEffect(() => {
    // Perform asynchronous operation to fetch session
    const fetchSession = async () => {
      try {
        const result = await getServerSession(authOptions);
        setSession(result);
      } catch (error) {
        // Handle any errors that occur during session retrieval
        console.error('Error fetching session:', error);
      }
    };

    // Call the async function to fetch session
    fetchSession();
  }, []);

  return (
    <html lang="en">
      <body className={`${montserrat.className} bg-primary min-h-screen min-w-screen  bg-white`}>
        <Provider session={session}>
          <CustomNavbar />
          <main className='h-screen bg-white'>
            {children}
          </main>
          <CustomFooter />
        </Provider>
      </body>
    </html>
  );
}