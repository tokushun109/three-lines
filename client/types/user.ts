export interface IUser {
    loginId: string
}

export interface ILoginForm extends IUser {
    password: string
}