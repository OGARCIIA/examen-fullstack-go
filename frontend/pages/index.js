import React, { useEffect, useState } from 'react';
import axios from 'axios';
import Layout from '../components/Layout';
import ProductList from '../components/ProductList';

const IndexPage = () => {
  const [token, setToken] = useState('');

  useEffect(() => {
    const getToken = async () => {
      try {
        const res = await axios.get('http://localhost:8080/generate-token');
        setToken(res.data.token);
      } catch (err) {
        console.error(err);
      }
    };

    getToken();
  }, []);

  if (!token) return <div>Generando token...</div>;

  return (
    <Layout>
      <h1>Productos</h1>
      <ProductList token={token} />
    </Layout>
  );
};

export default IndexPage;
