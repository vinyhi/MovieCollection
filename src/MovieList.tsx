import React, { useState, useEffect } from 'react';

interface IMovie {
  id: string;
  title: string;
  director: string;
  genre: string;
  releaseDate: string;
  review: string;
  rating: number;
}

const MoviesList: React.FC = () => {
  const [movies, setMovies] = useState<IMovie[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchMovies = async () => {
      setIsLoading(true);
      try {
        const response = await fetch(`${process.env.REACT_APP_BACKEND_URL}/movies`);
        if (!response.ok) {
          throw new Error('Something went wrong!');
        }
        const data: IMovie[] = await response.json();
        setMovies(data);
      } catch (err) {
        setError(err.message);
      } finally {
        setIsLoading(false);
      }
    };

    fetchMovies();
  }, []);

  const handleDelete = async (id: string) => {
    try {
      await fetch(`${process.env.REACT_APP_BACKEND_URL}/movies/${id}`, {
        method: 'DELETE',
      });
      setMovies(movies.filter(movie => movie.id !== id));
    } catch (err) {
      console.error(err);
    }
  };

  const handleUpdate = (id: string) => {
    console.log(`Update movie with id: ${id}`);
  };

  return (
    <div>
      {isLoading ? (
        <p>Loading...</p>
      ) : error ? (
        <p>{error}</p>
      ) : (
        <table>
          <thead>
            <tr>
              <th>Title</th>
              <th>Director</th>
              <th>Genre</th>
              <th>Release Date</th>
              <th>Review</th>
              <th>Rating</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {movies.map((movie) => (
              <tr key={movie.id}>
                <td>{movie.title}</td>
                <td>{movie.director}</td>
                <td>{movie.genre}</td>
                <td>{movie.releaseDate}</td>
                <td>{movie.review}</td>
                <td>{movie.rating}</td>
                <td>
                  <button onClick={() => handleUpdate(movie.id)}>Update</button>
                  <button onClick={() => handleDelete(movie.id)}>Delete</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
};

export default Movies, List;