import { useLocation } from "react-router-dom"
import { Link } from "react-router-dom"

import { useState } from "react"
import { useEffect } from "react"
import { MovieList } from "../../api"
import { DefaultService } from "../../api"

import MovieWidget from "./MovieWidget"

import leftButton from "../../assets/leftbutton.svg"
import rightButton from "../../assets/rightbutton.svg"

interface Props {
    query: string, 
    page: string,
}

const MovieResults = ({query, page}: Props) => {
    
    const currentPage = Number(page)

    const [searchResults, setSearchResults] = useState<MovieList>({page: 0,
        totalPages: 0,
        results: []});

    useEffect(() => {
        DefaultService.searchMovie(query, page )
        .then((result) => {
            setSearchResults(result)
        })
        .catch((error) => {
            console.error('Error: ', error)
        })
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [useLocation()])

    const movies = (searchResults.results.map(MovieWidget))

    return (
        <div className="flex justify-center">
            <div className="w-12">
                {currentPage > 1 ? 
                    <Link to={"/search?query=" + query + "&page=" + (currentPage - 1)} className="fixed h-12 w-12 top-1/2">
                        <img src={leftButton} />
                    </Link>
                : <></>}
            </div>
            <div className="static">
                <div className="grid 2xl:grid-cols-5 lg:grid-cols-3 grid-cols-1 gap-x-2 gap-y-12 pb-12">
                    {...movies}
                </div>
            </div>
            <div className="w-12">
                {currentPage < searchResults.totalPages ? 
                    <Link to={"/search?query=" + query + "&page=" + (currentPage + 1)} className="fixed h-12 w-12 top-1/2">
                        <img className="fill-gray-400" src={rightButton} />
                    </Link> 
                : <></>}
            </div>
        </div>
    )
}

export default MovieResults;