import NextAuth from "next-auth"
import type { AuthOptions } from "next-auth"
import CredentialsProvider from "next-auth/providers/credentials"
import GitHubProvider from "next-auth/providers/github";
import { authenticate } from "@/services/authService"


export const authOptions: AuthOptions = {
    providers: [
        GitHubProvider({
            clientId:process.env.NEXT_PUBLIC_AUTH_GITHUB_ID,
            clientSecret:process.env.NEXT_PUBLIC_AUTH_GITHUB_SECRET
        }),
        CredentialsProvider({
            name: 'Email',
            credentials: {
                email: { label: "Email", type: "text" },
                password: { label: "Password", type: "password" }
            },
            async authorize(credentials, req) {
                if (typeof credentials !== "undefined") {
                    const res = await authenticate(credentials.email, credentials.password)
                    if (typeof res !== "undefined") {
                        return { ...res.user, apiToken: res.token }
                    } else {
                        return null
                    }
                } else {
                    return null
                }
            }
        })
    ],
    session: { strategy: "jwt" }
}

const handler = NextAuth(authOptions)

export { handler as GET, handler as POST }
