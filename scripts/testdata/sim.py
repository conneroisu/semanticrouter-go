
import numpy as np


def similarity_matrix(xq, index) -> np.ndarray:
    xq_norm = np.linalg.norm(xq)
    index_norm = np.linalg.norm(index)
    dot_product = np.dot(xq, index)
    return dot_product / (xq_norm * index_norm)


xq = np.array([1, 2, 3])
index = np.array([4, 5, 6])

similarity = similarity_matrix(xq, index)
print(f"Similarity: {similarity}")
