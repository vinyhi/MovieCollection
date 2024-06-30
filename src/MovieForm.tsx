import React, { useState } from 'react';

interface MovieFormState {
  title: string;
  director: string;
  genre: string;
  releaseDate: string;
  review: string;
  rating: number;
}

const initialFormState: MovieFormState = {
  title: '',
  director: '',
  genre: '',
  releaseDate: '',
  review: '',
  rating: 0,
};

const AddMovieForm: React.FC = () => {
  const [formData, setFormData] = useState<MovieFormState>(initialFormState);
  const [message, setMessage] = useState<string>('');

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      await sendMovieData(formData);
      setMessage('Movie added successfully');
      setFormData(initialFormState);
    } catch (error) {
      setMessage('Error adding movie');
    }
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="title">Title:</label>
          <input
            name="title"
            type="text"
            value={formData.title}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor="director">Director:</label>
          <input
            name="director"
            type="text"
            value={formData.director}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor="genre">Genre:</label>
          <input
            name="genre"
            type="text"
            value={formData.genre}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor="releaseDate">Release Date:</label>
          <input
            name="releaseDate"
            type="date"
            value={formData.releaseDate}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor="review">Review:</label>
          <textarea
            name="review"
            value={formData.review}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor="rating">Rating:</label>
          <input
            name="rating"
            type="number"
            value={formData.rating}
            onChange={handleChange}
            min="0"
            max="10"
          />
        </div>

        <button type="submit">Submit</button>
      </form>
      {message && <div>{message}</div>}
    </div>
  );
};

async function sendMovieData(movieData: MovieFormState): Promise<void> {
  console.log(movieData);
}

export default AddMovieForm;