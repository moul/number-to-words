require "language/go"

class NumberToWords < Formula
  desc "number-to-words: convert number into words"
  homepage "https://github.com/moul/number-to-words"
  url "https://github.com/moul/number-to-words/archive/v0.1.0.tar.gz"
  sha256 "0d4aaf616f619daf0b4d17d281c19dadbebc2b3294a88783101b9e474cde9600"

  head "https://github.com/moul/number-to-words.git"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    ENV["GOBIN"] = buildpath
    (buildpath/"src/github.com/moul/number-to-words").install Dir["*"]

    system "go", "build", "-o", "#{bin}/number-to-words", "-v", "github.com/moul/number-to-words/cmd/number-to-words/"
  end
end
