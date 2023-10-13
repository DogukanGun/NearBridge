export {}
declare global{
    namespace NodeJS{
        interface ProcessEnv{
            NEXT_PUBLIC_AUTH_GITHUB_ID: string
            NEXT_PUBLIC_AUTH_GITHUB_SECRET:string
            NEXT_PUBLIC_SECRET:string
            NEXT_PUBLIC_NEXTAUTH_URL:string
            NEXT_PUBLIC_NEXTAUTH_SECRET:string
            NEXT_PUBLIC_API_URL:string
        }
    }
}