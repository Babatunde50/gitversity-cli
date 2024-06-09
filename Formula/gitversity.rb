# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Gitversity < Formula
  desc ""
  homepage ""
  version "0.1.21"

  on_macos do
    on_intel do
      url "https://github.com/Babatunde50/gitversity-cli/releases/download/v0.1.21/gitversity-cli_Darwin_x86_64.tar.gz"
      sha256 "caeb3251b097609ca04b5bd3236666d784056ff7539eedff619d5f10deb69f38"

      def install
        bin.install "gitversity-cli"
        bin.install_symlink bin/"gitversity-cli" => "gitversity"
      end
    end
    on_arm do
      url "https://github.com/Babatunde50/gitversity-cli/releases/download/v0.1.21/gitversity-cli_Darwin_arm64.tar.gz"
      sha256 "f089d196481398835cd4ef2227fdb295e1cebb0d164103d7470397fab0ed1726"

      def install
        bin.install "gitversity-cli"
        bin.install_symlink bin/"gitversity-cli" => "gitversity"
      end
    end
  end

  on_linux do
    on_intel do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/Babatunde50/gitversity-cli/releases/download/v0.1.21/gitversity-cli_Linux_x86_64.tar.gz"
        sha256 "e350cc0903077f8f89d5b861c521949c9f11d3dfbe14e5f490a35bf7348f938c"

        def install
          bin.install "gitversity-cli"
          bin.install_symlink bin/"gitversity-cli" => "gitversity"
        end
      end
    end
    on_arm do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/Babatunde50/gitversity-cli/releases/download/v0.1.21/gitversity-cli_Linux_arm64.tar.gz"
        sha256 "8f59a9442a0cad36bee4cb160641b342f1363440fdeba1257460759648aeffcd"

        def install
          bin.install "gitversity-cli"
          bin.install_symlink bin/"gitversity-cli" => "gitversity"
        end
      end
    end
  end
end
