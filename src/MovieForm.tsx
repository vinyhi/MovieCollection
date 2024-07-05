import React, { useState } from 'react';

interface MovieFormState {
  title: string;
  director: string;
  genre: string;
  releaseDate: string;
  review: string;
  rating: number;
  image?: File; // Optional image attribute
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
    if (name === "image") {
      // Handling file input for image
      const file = (e.target as HTMLInputElement).files![0]; // Assuming single file upload
      setFormData({ ...formData, image: file });
    } else {
      setFormData({ ...formData, [name]: value });
    }
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      // Preparing FormData to include files in the request
      const movieData = new FormData();
      Object.keys(formData).forEach((key) => {
        if (key === "image" && formData.image) {
          movieData.append(key, formData.image);
        } else {
          movieData.append(key, formData[key].toString());
        }
      });

      await sendMovieData(movieData);
      setMessage('Movie added successfully');
      setFormData(initialFormState); // Resetting form to initial state
    } catch (error) {
      setMessage('Error adding movie');
    }
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        {/* Form inputs for movie details */}
        <div>
          <label htmlFor="image">Movie Image:</label>
          <input
            name="image"
            type="file"
            onChange={handleChange}
          />
        </div>

        <button type="submit">Submit</button>
      </form>
      {message && <div>{message}</div>}
    </div>
  );
};

async function sendMovieData(movieData: FormData): Promise<void> {
  console.log('Sending movie data...', movieData);
  // Placeholder for API call. Use fetch or axios with 'Content-Type': 'multipart/form-data'
}

export default AddMovie resend;