package bc

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// Take the data from the Block

// Create a counter (nonce) which starts at 0

// Create a hash of the data plus the counter

// Check the hash to see if it meets a set of requirments (difficulty)

/**
 * Requirments:
 * The First few bytes must contain 0s
 */

const Difficulty = 12 // generally, slowly increasing in real Blockchain algorithm

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) { // main computaional function
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

/*
https://www.youtube.com/watch?v=bBC-nXj3Ng4
But how does bitcoin actually work?
조회수 8,693,172회•2017. 7. 8.

21만

2.7천

공유

저장


3Blue1Brown
구독자 366만명

The math behind cryptocurrencies.
Help fund future projects: https://www.patreon.com/3blue1brown​
An equally valuable form of support is to simply share some of the videos.
Special thanks to these supporters: http://3b1b.co/btc-thanks​
This video was also funded with help from Protocol Labs: https://protocol.ai/join/​

Some people have asked if this channel accepts contributions in cryptocurrency form.  As a matter of fact, it does:
http://3b1b.co/crypto​

2^256 video: https://youtu.be/S9JGmA5_unY​

Music by Vincent Rubinetti: https://soundcloud.com/vincerubinetti...​

Here are a few other resources I'd recommend:

Original Bitcoin paper: https://bitcoin.org/bitcoin.pdf​

Block explorer: https://blockexplorer.com/​

Blog post by Michael Nielsen: https://goo.gl/BW1RV3​
(This is particularly good for understanding the details of what transactions look like, which is something this video did not cover)

Video by CuriousInventor: https://youtu.be/Lx9zgZCMqXE​

Video by Anders Brownworth: https://youtu.be/_160oMzblY8​

Ethereum white paper: https://goo.gl/XXZddT​

------------------
Animations largely made using manim, a scrappy open source python library.  https://github.com/3b1b/manim​

If you want to check it out, I feel compelled to warn you that it's not the most well-documented tool, and has many other quirks you might expect in a library someone wrote with only their own use in mind.

Music by Vincent Rubinetti.
Download the music on Bandcamp:
https://vincerubinetti.bandcamp.com/a...​

Stream the music on Spotify:
https://open.spotify.com/album/1dVyjw...​

If you want to contribute translated subtitles or to help review those that have already been made by others and need approval, you can click the gear icon in the video and go to subtitles/cc, then "add subtitles/cc".  I really appreciate those who do this, as it helps make the lessons accessible to more people.
------------------

3blue1brown is a channel about animating math, in all senses of the word animate.  And you know the drill with YouTube, if you want to stay posted on new videos, subscribe, and click the bell to receive notifications (if you're into that).

If you are new to this channel and want to see more, a good place to start is this playlist: http://3b1b.co/recommended​

Various social media stuffs:
Website: https://www.3blue1brown.com​
Twitter: https://twitter.com/3Blue1Brown​
Patreon: https://patreon.com/3blue1brown​
Facebook: https://www.facebook.com/3blue1brown​
Reddit: https://www.reddit.com/r/3Blue1Brown

*/
