# FbHash-go: A New Similarity Hashing Scheme for Digital Forensics
golang implementation of Chang, et al's FbHash algorithms for generating similarity preserving cryptographic hashes


## Citations:
Donghoon Chang, Mohona Ghosh, Somitra Kumar Sanadhya, Monika Singh, Douglas R. White\
FbHash: A New Similarity Hashing Scheme for Digital Forensics\
Digital Investigation\
Volume 29, Supplement\
2019\
Pages S113-S123\
ISSN 1742-2876\
[https://doi.org/10.1016/j.diin.2019.04.006.](http://www.sciencedirect.com/science/article/pii/S1742287619301550)\
Abstract: With the rapid growth of the World Wide Web and Internet of Things, a huge amount of digital data is being produced every day. Digital forensics investigators face an uphill task when they have to manually screen through and examine tons of such data during a crime investigation. To expedite this process, several automated techniques have been proposed and are being used in practice. Among which tools based on Approximate Matching algorithms have gained prominence, e.g., ssdeep, sdhash, mvHash etc. These tools produce hash signatures for all the files to be examined, compute a similarity score and then compare it with a known reference set to filter out known good as well as bad files. In this way, exact as well as similar matches can be screened out. However, all of these schemes have been shown to be prone to active adversary attack, whereby an attacker, by making feasible changes in the content of the file, intelligently modifies the final hash signature produced to evade detection. Thus, an alternate hashing scheme is required which can resist this attack. In this work, we propose a new Approximate Matching scheme termed as - FbHash. We show that our scheme is secure against active attacks and detects similarity with 98% accuracy. We also provide a detailed comparative analysis with other existing schemes and show that our scheme has a 28% higher accuracy rate than other schemes for uncompressed file format (e.g., text files) and 50% higher accuracy rate for compressed file format (e.g., docx etc.). Our proposed scheme is able to correlate a fragment as small as 1% to the source file with 100% detection rate and able to detect commonality as small as 1% between two documents with appropriate similarity score. Further, our scheme also produces the least false negatives in comparison to other schemes.\
Keywords: *Data fingerprinting; Similarity digests; Fuzzy hashing; TF-IDF; Cosine-similarity*

