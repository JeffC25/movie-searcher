/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Cast } from '../models/Cast';
import type { MovieDetails } from '../models/MovieDetails';
import type { MovieList } from '../models/MovieList';
import type { ReviewList } from '../models/ReviewList';
import type { VideoList } from '../models/VideoList';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class DefaultService {

    /**
     * Get currently playing movies
     * @param page page number
     * @returns MovieList OK
     * @throws ApiError
     */
    public static getNowPlaying(
        page: string,
    ): CancelablePromise<MovieList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/nowplaying',
            query: {
                'page': page,
            },
            errors: {
                400: `Bad Request`,
                418: `I'm a teapot`,
                500: `Internal Server Error`,
                502: `Bad Gateway`,
            },
        });
    }

    /**
     * Get top rated movies
     * @param page page number
     * @returns MovieList OK
     * @throws ApiError
     */
    public static getTopRated(
        page: string,
    ): CancelablePromise<MovieList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/toprated',
            query: {
                'page': page,
            },
            errors: {
                400: `Bad Request`,
                418: `I'm a teapot`,
                500: `Internal Server Error`,
                502: `Bad Gateway`,
            },
        });
    }

    /**
     * Get popular movies
     * @param page page number
     * @returns MovieList OK
     * @throws ApiError
     */
    public static getPopular(
        page: string,
    ): CancelablePromise<MovieList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/popular',
            query: {
                'page': page,
            },
            errors: {
                400: `Bad Request`,
                418: `I'm a teapot`,
                500: `Internal Server Error`,
                502: `Bad Gateway`,
            },
        });
    }

    /**
     * Get upcoming movies
     * @param page page number
     * @returns MovieList OK
     * @throws ApiError
     */
    public static getUpcoming(
        page: string,
    ): CancelablePromise<MovieList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/upcoming',
            query: {
                'page': page,
            },
            errors: {
                400: `Bad Request`,
                418: `I'm a teapot`,
                500: `Internal Server Error`,
                502: `Bad Gateway`,
            },
        });
    }

    /**
     * Search for movie
     * @param queryString Query string
     * @param page page number
     * @returns MovieList OK
     * @throws ApiError
     */
    public static searchMovie(
        queryString: string,
        page: string,
    ): CancelablePromise<MovieList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/search',
            query: {
                'queryString': queryString,
                'page': page,
            },
            errors: {
                400: `Bad Request`,
                418: `I'm a teapot`,
                500: `Internal Server Error`,
                502: `Bad Gateway`,
            },
        });
    }

    /**
     * Get movie details by ID
     * @param movieId Numeric ID of movie to get
     * @returns MovieDetails OK
     * @throws ApiError
     */
    public static getMovieDetails(
        movieId: string,
    ): CancelablePromise<MovieDetails> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/details/{movieId}',
            path: {
                'movieId': movieId,
            },
            errors: {
                400: `Bad Request`,
                418: `I'm a teapot`,
                500: `Internal Server Error`,
                502: `Bad Gateway`,
            },
        });
    }

    /**
     * Get movie reviews by ID
     * @param movieId Numeric ID of movie to get
     * @param page page number
     * @returns ReviewList OK
     * @throws ApiError
     */
    public static getMovieReviews(
        movieId: string,
        page: string,
    ): CancelablePromise<ReviewList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/reviews/{movieId}',
            path: {
                'movieId': movieId,
            },
            query: {
                'page': page,
            },
            errors: {
                400: `Bad Request`,
                418: `I'm a teapot`,
                500: `Internal Server Error`,
                502: `Bad Gateway`,
            },
        });
    }

    /**
     * Get movie videos by ID
     * @param movieId Numeric ID of movie to get
     * @returns VideoList OK
     * @throws ApiError
     */
    public static getMovieVideos(
        movieId: string,
    ): CancelablePromise<VideoList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/videos/{movieId}',
            path: {
                'movieId': movieId,
            },
            errors: {
                400: `Bad Request`,
                418: `I'm a teapot`,
                500: `Internal Server Error`,
                502: `Bad Gateway`,
            },
        });
    }

    /**
     * Get movie cast by ID
     * @param movieId Numeric ID of movie to get
     * @returns Cast OK
     * @throws ApiError
     */
    public static getMovieCast(
        movieId: string,
    ): CancelablePromise<Cast> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/cast/{movieId}',
            path: {
                'movieId': movieId,
            },
            errors: {
                400: `Bad Request`,
                418: `I'm a teapot`,
                500: `Internal Server Error`,
                502: `Bad Gateway`,
            },
        });
    }

}
