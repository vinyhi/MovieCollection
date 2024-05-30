import React, { useState } from 'react';
import MovieForm from './MovieForm';
import MovieList from './MovieList';

type Movie = {
  id: number;
  title: string;
  year: number;
};

const App: React.FC = () => {
  const [movies, setMovies] = useState<Movie[]>([]);

  const addMovie = (newMovie: Omit<Movie, 'id'>) => {
    const movieToAdd: Movie = {
      id: movies.length + 1,
      ...newMovie,
    };
    setMovies([...movies, movieToAdd]);
  };

  return (
    <div>
      <header>
        <h1>MovieCollection</h1>
      </header>
      <MovieForm onAddMovie={addMovie} />
      <MovieList movies={movies} />
    </div>
  );
};

export default App;