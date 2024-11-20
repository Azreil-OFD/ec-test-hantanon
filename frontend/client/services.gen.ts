// This file is auto-generated by @hey-api/openapi-ts

import { createClient, createConfig, type Options } from '@hey-api/client-fetch';
import type { LoginData, LoginError, LoginResponse, RegisterData, RegisterError, RegisterResponse, GetProfileError, GetProfileResponse, SendFriendRequestData, SendFriendRequestError, SendFriendRequestResponse, AcceptFriendRequestData, AcceptFriendRequestError, AcceptFriendRequestResponse, DeclineFriendRequestData, DeclineFriendRequestError, DeclineFriendRequestResponse, RemoveFriendData, RemoveFriendError, RemoveFriendResponse, SearchUserData, SearchUserError, SearchUserResponse } from './types.gen';

export const client = createClient(createConfig());

/**
 * Авторизация пользователя
 * Получение JWT токена для пользователя после успешной авторизации.
 */
export const login = <ThrowOnError extends boolean = false>(options: Options<LoginData, ThrowOnError>) => {
    return (options?.client ?? client).post<LoginResponse, LoginError, ThrowOnError>({
        ...options,
        url: '/api/auth'
    });
};

/**
 * Регистрация пользователя
 * Регистрирует нового пользователя в системе, принимая логин, пароль, email и полное имя.
 */
export const register = <ThrowOnError extends boolean = false>(options: Options<RegisterData, ThrowOnError>) => {
    return (options?.client ?? client).post<RegisterResponse, RegisterError, ThrowOnError>({
        ...options,
        url: '/api/register'
    });
};

/**
 * Получить профиль текущего пользователя
 * Возвращает профиль пользователя, извлекаемый из контекста JWT токена.
 */
export const getProfile = <ThrowOnError extends boolean = false>(options?: Options<unknown, ThrowOnError>) => {
    return (options?.client ?? client).get<GetProfileResponse, GetProfileError, ThrowOnError>({
        ...options,
        url: '/api/profile'
    });
};

/**
 * Отправить запрос на добавление в друзья
 * Отправляет запрос на добавление в друзья указанному пользователю.
 */
export const sendFriendRequest = <ThrowOnError extends boolean = false>(options: Options<SendFriendRequestData, ThrowOnError>) => {
    return (options?.client ?? client).post<SendFriendRequestResponse, SendFriendRequestError, ThrowOnError>({
        ...options,
        url: '/api/friends/request'
    });
};

/**
 * Принять запрос на добавление в друзья
 * Принять запрос на добавление в друзья от другого пользователя.
 */
export const acceptFriendRequest = <ThrowOnError extends boolean = false>(options: Options<AcceptFriendRequestData, ThrowOnError>) => {
    return (options?.client ?? client).post<AcceptFriendRequestResponse, AcceptFriendRequestError, ThrowOnError>({
        ...options,
        url: '/api/friends/accept'
    });
};

/**
 * Отклонить запрос на добавление в друзья
 * Отклонить запрос на добавление в друзья от другого пользователя.
 */
export const declineFriendRequest = <ThrowOnError extends boolean = false>(options: Options<DeclineFriendRequestData, ThrowOnError>) => {
    return (options?.client ?? client).post<DeclineFriendRequestResponse, DeclineFriendRequestError, ThrowOnError>({
        ...options,
        url: '/api/friends/decline'
    });
};

/**
 * Удалить друга
 * Удаляет друга из списка друзей.
 */
export const removeFriend = <ThrowOnError extends boolean = false>(options: Options<RemoveFriendData, ThrowOnError>) => {
    return (options?.client ?? client).post<RemoveFriendResponse, RemoveFriendError, ThrowOnError>({
        ...options,
        url: '/api/friends/remove'
    });
};

/**
 * Поиск пользователей
 * Ищет пользователей по логину или полному имени с пагинацией.
 */
export const searchUser = <ThrowOnError extends boolean = false>(options: Options<SearchUserData, ThrowOnError>) => {
    return (options?.client ?? client).get<SearchUserResponse, SearchUserError, ThrowOnError>({
        ...options,
        url: '/api/search'
    });
};