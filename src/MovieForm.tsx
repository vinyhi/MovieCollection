import React, { useState } from 'react';

interface MovieFormState {
  title: string;
  director: string;
  genre: string;
  releaseDate: string;
  review: string;
  rating: number;
  image?: File; // Added image to the movie form state
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
    if(name === "image"){ // Check if the input type is file for image
      const file = (e.target as HTMLInputElement).files![0]; // Assuming single file upload
      setFormData({ ...formData, image: file });
    }else{
      setFormData({ ...formData, [name]: value });
    }
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      // Assuming sendMovieData can now handle FormData to include files
      const movieData = new FormData();
      Object.keys(formData).forEach((key) => {
        if(key === "image"){
          if(formData.image) movieData.append(key, formData.image);
        }else{
          movieData.append(key, formData[key].toString());
        }
      });

      await sendMovieData(movieData);
      setMessage('Movie added successfully');
      setFormData(initialFormState);
    } catch (error) {
      setMessage('Error adding movie');
    }
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        {/* Input fields remain unchanged, except the addition of File input for image */}
        {/* ...previous input scores */}

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
  // Assuming this will be replaced with an actual API call that supports FormData.
  // For API calls, you'd generally use fetch or axios and set 'Content-Type': 'multipart/form-data' in headers.
}

export default Add 