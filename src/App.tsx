import React, { useState } from 'react';
import MovieForm from './MovieForm';
import MovieList from './MovieList';

type Movie = {
  id: number;
  title: string;
  year: number;
};

type ErrorState = {
  isError: boolean;
  message: string;
};

const App: React.FC = () => {
  const [movies, setMovies] = useState<Movie[]>([]);
  const [error, setError] = useState<ErrorState>({ isError: false, message: '' });

  const validateMovie = (newMovie: Omit<Movie, 'id'>): boolean => {
    if (!newMovie.title.trim()) {
      setError({ isError: true, message: 'Movie title cannot be empty.' });
      return false;
    }
    if (newMovie.year < 1800 || newMovie.year > new Date().getFullYear()) {
      setError({ isError: true, message: 'Year is not valid. It should be between 1800 and the current year.' });
      return false;
    }
    setError({ isError: false, message: '' });
    return true;
  }

  const addMovie = (newMovie: Omit<Movie, 'id'>) => {
    if (!validateMovie(newMovie)) {
      return;
    }

    const movieToAdd: Movie = {
      id: movies.length > 0 ? Math.max(...movies.map(movie => movie.id)) + 1 : 1,
      ...newMovie,
    };
    setMovies([...movies, movieToAdd]);
  };

  return (
    <div>
      <header>
        <h1>MovieCollection</h1>
      </header>
      {error.isError && <p style={{ color: 'red' }}>{error.message}</p>}
      <MovieForm onAddMovie={addMovie} />
      <MovieList movies={movies} />
    </div>
  );
};

export default App;